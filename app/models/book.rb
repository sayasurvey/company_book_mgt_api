class Book < ApplicationRecord
  belongs_to :user
  has_many :borrowed_records

  # 基本的なスコープ
  scope :available, -> { where(status: 'available') }
  scope :borrowed, -> { where(status: 'borrowed') }
  
  # 引数を受け取るスコープ
  scope :by_category, ->(category) { where(category: category) }
  scope :published_after, ->(date) { where('published_date > ?', date) }
  
  # 関連テーブルを含むスコープ
  scope :with_active_users, -> { joins(:user).merge(User.active) }
  
  # 複雑な条件を含むスコープ
  scope :recently_borrowed, -> { 
    joins(:borrowed_records)
      .where('borrowed_records.created_at > ?', 1.month.ago)
      .distinct
  }
  
  # デフォルト値を持つスコープ
  scope :by_title, ->(query = '') { 
    where('title LIKE ?', "%#{query}%") if query.present?
  }
  
  # 複数のスコープを組み合わせた例
  scope :available_in_category, ->(category) {
    available.by_category(category)
  }
end 