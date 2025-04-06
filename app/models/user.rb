class User < ApplicationRecord
  # リレーションの定義
  has_many :books
  belongs_to :department

  # スコープの定義
  scope :active, -> { where(active: true) }
  scope :created_last_week, -> { where(created_at: 1.week.ago..Time.current) }
  scope :by_name, ->(name) { where("name LIKE ?", "%#{name}%") }
  scope :with_books, -> { joins(:books).distinct }
  
  # 引数を受け取るスコープ
  scope :created_before, ->(date) { where("created_at < ?", date) }
  
  # 複数の条件を組み合わせたスコープ
  scope :active_admins, -> { active.where(role: 'admin') }
  
  # 他のスコープを含むスコープ
  scope :recent_active_users, -> { active.created_last_week }
end 